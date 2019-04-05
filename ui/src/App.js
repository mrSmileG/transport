import React, { Component } from 'react';
import './App.css';
import Layout from './components/Layout';
import NewTaskModal from './components/NewTaskModal';
import Matrix from './components/Matrix';
import { Card, Button } from 'semantic-ui-react';

class App extends Component {
  state = {
    providersCount: '',
    consumersCount: '',
    providers: [],
    consumers: [],
    rows: [],
    columns: [],
    columnsCons: [],
    error: null,
    isLoaded: false,
    items: []
  };

  initialState = {
    providersCount: '',
    consumersCount: '',
    providers: [],
    consumers: [],
    rows: [],
    columns: [],
    columnsCons: []
  };


  updateDate = (providers, consumers) => {
    this.setState({ providersCount: providers });
    this.setState({ consumersCount: consumers });
    this.setState({
      providers: [],
      consumers: [],
      rows: [],
      columns: [],
      columnsCons: []
    });
  }

  renderMatrix() {
    const {
      providersCount,
      consumersCount,
      providers,
      consumers,
      rows,
      columns,
      columnsCons
    } = this.state;

    const items = [
      {
          header: 'Providers vector data',
          meta: 'P1 - P(n) - Providers. Enter the resources count of every provider',
          description: (
            <Matrix providersCount={providersCount} rows={providers} columns={columns} />
          ),
          fluid: true,
          style: {padding: '20px 40px'}
      },
      {
          header: 'Consumers vector data',
          meta: 'C1 - C(n) - Consumers. Enter the resources count to need for every consumer',
          description: (
            <Matrix providersCount={consumersCount} rows={consumers} columns={columnsCons} />
          ),
          fluid: true,
          style: {padding: '20px 40px'}
      },
      {
          header: 'Matrix of prices',
          meta: 'Enter the delivery prices from every provider to every consumer',
          description: (
            <Matrix providersCount={providersCount} consumersCount={consumersCount} rows={rows} columns={columns} />
          ),
          fluid: true,
          style: {padding: '20px 40px'}
      }
  ];

    if (providersCount > 0 && consumersCount > 0) {
      for (let i = 0; i < providersCount; i++){
        columns.push({ key: i, name: "P"+(i+1), editable: true })
      }
      for (let i = 0; i < consumersCount; i++){
          rows.push({i: 0});
          columnsCons.push({ key: i, name: "C"+(i+1), editable: true })
      }
      providers.push({0: ''});
      consumers.push({0: ''});
      return <div> 
          <Card.Group style={{ marginTop: '70px', textAlign: 'left' }} itemsPerRow="1" items={items} />
        </div>
    }
  }

  solve = () => {
    fetch('https://warm-ridge-60909.herokuapp.com/transport/', {
        method: 'POST',
        mode: 'no-cors',
        body: JSON.stringify({ 
          providers: [12, 40, 33], 
          consumers: [20, 30, 10], 
          prices: [[3, 5, 7], [2, 4, 6], [9, 1, 8]]
        })
      })
      .then(function(res){ return res; })
      .then(function(data){ console.log( data.json() )})
  }

  renderSolution() {
    const { items } = this.state;
      return (
        
        JSON.stringify(items)
        // <ul>
        //   {items.map(item => (
        //     <li key={item.name}>
        //       {item.name} {item.price}
        //     </li>
        //   ))}
        // </ul>
      );
    
  }

  render() {
    return (
        <Layout>
          <NewTaskModal updateData={this.updateDate} />
          {this.renderMatrix()}
          <Button 
              positive 
              onClick={this.solve} 
              icon="cancel"
              content="Get Solution"
              style={{marginRight: '10px'}}
          />
          {this.renderSolution()}
        </Layout>
    );
  }
}

export default App;
