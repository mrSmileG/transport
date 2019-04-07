import React, { Component } from 'react';
import './App.css';
import Layout from './components/Layout';
import NewTaskModal from './components/NewTaskModal';
import Matrix from './components/Matrix';
import { Card, Button, Message, Form } from 'semantic-ui-react';
import SolutionTable from './components/SolutionTable';

class App extends Component {
  state = {
    providersCount: '',
    consumersCount: '',
    providers: [],
    consumers: [],
    rows: [],
    columns: [],
    columnsCons: [],
    errorMessage: '',
    isLoaded: false,
    loading: false,
    items: [],
    provVector: [],
    consVector: [],
    prices: []
  };

  updateProvider = (rows) => {
    this.setState({
      provVector: rows
    });
  }

  updateConsumer = (rows) => {
    this.setState({
      consVector: rows
    });
  }

  updatePrices = (rows) => {
    this.setState({
      prices: rows
    });
  }

  updateData = (providers, consumers) => {
    this.setState({ providersCount: providers });
    this.setState({ consumersCount: consumers });

    this.setState({
      columns: [],
      rows: [],
      providers: [],
      consumers: [],
      columnsCons: []
    })
  }

  renderMatrix() {
    const {
      providersCount,
      consumersCount,
      providers,
      consumers,
      rows,
      columns,
      columnsCons,
      loading
    } = this.state;

    const items = [
      {
          header: 'Providers vector data',
          meta: 'P1 - P(n) - Providers. Enter the resources count of every provider',
          description: (
            <Matrix providersCount={providersCount} rows={providers} columns={columns} update={this.updateProvider} />
          ),
          fluid: true,
          style: {padding: '20px 40px'}
      },
      {
          header: 'Consumers vector data',
          meta: 'C1 - C(n) - Consumers. Enter the resources count to need for every consumer',
          description: (
            <Matrix providersCount={consumersCount} rows={consumers} columns={columnsCons} update={this.updateConsumer} />
          ),
          fluid: true,
          style: {padding: '20px 40px'}
      },
      {
          header: 'Matrix of prices',
          meta: 'Enter the delivery prices from every provider to every consumer',
          description: (
            <Matrix providersCount={providersCount} consumersCount={consumersCount} rows={rows} columns={columns} update={this.updatePrices} />
          ),
          fluid: true,
          style: {padding: '20px 40px'}
      }
    ];

    

    if (providersCount > 0 && consumersCount > 0) {

      for (let i = 0; i < providersCount; i++){
        columns.pop();
      }
      for (let i = 0; i < consumersCount; i++){
          rows.pop();
          columnsCons.pop();
      }
      providers.pop();
      consumers.pop();


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
          <Button 
              positive 
              floated="right"
              onClick={this.solve} 
              icon="check"
              content="Get Solution"
              loading={loading} 
              disabled={loading} 
              style={{margin: '20px  0'}}
          />
          {this.renderSolution()}
        </div>
    }
  }

  solve = () => {

    this.setState({
      errorMessage: '',
      loading: true
    });

    if (this.state.provVector[0] === undefined) return;
    fetch('https://mighty-shelf-82507.herokuapp.com/', {
        method: 'POST',
        body: JSON.stringify({ 
          // // providers: [12, 40, 33],
          // //     consumers: [20, 30, 10],
          // //     prices: [[3, 5, 7], [2, 4, 6], [9, 1, 8]]
          // providers: [12, 40, 33], 
          // consumers: [20, 30, 10], 
          // prices: [[3, 5, 7], [2, 4, 6], [9, 1, 8]]


            providers: Object.values(
              this.state.provVector[0])
              .map(item => (parseInt(item, 10))),
            consumers: Object.values(
              this.state.consVector[0])
              .map(item => (parseInt(item, 10))),
            prices: Object.values(
              this.state.prices)
              .map(items => Object.values(items).map(item => (parseInt(item, 10))))
        })
      })
      .then(res => res.json())
      .then((result) => {
          this.setState({
            isLoaded: true,
            items: result
          });
        },
        (error) => {
          this.setState({
            isLoaded: true,
            errorMessage: error.message
          });
        }
      )

    this.setState({
        loading: false
    });
  }

  renderSolution() {
    const { items } = this.state;
      return (
        <SolutionTable data={items} />
      );
    
  }

  render() {
    return (
        <Layout>
           <Form  error={!!this.state.errorMessage}>
            <Message  error header="Oops!" content={this.state.errorMessage} />
          </Form>
          <NewTaskModal updateData={this.updateData} />
          {this.renderMatrix()}
        </Layout>
    );
  }
}

export default App;
