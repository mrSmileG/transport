import React, { Component } from 'react';

class Solver extends Component {
    state = {
        error: null,
        isLoaded: false,
        items: []
    }
    
      solve = () => {
        fetch('https://warm-ridge-60909.herokuapp.com/transport/', {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              "providers": [12, 40, 33],
              "consumers": [20, 30, 10],
              "prices": [[3, 5, 7], [2, 4, 6], [9, 1, 8]]
            })
          })
          .then(res => res.json())
          .then(
            (result) => {
              alert(result);
              this.setState({
                isLoaded: true,
                items: result
              });
            },                                                                                          
            (error) => {
              this.setState({
                isLoaded: true,
                error
              });
            }
          )
      }
    
      render() {
        const { error, isLoaded, items } = this.state;
        if (error) {
          return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
          return <div>Loading...</div>;
        } else {
          alert(items)
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
      }
}

export default Solver;