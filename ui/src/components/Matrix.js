import ReactDataGrid from "react-data-grid";
import React, { Component } from 'react';

class Matrix extends Component {
  static async getInitialProps(props) {
    const {
      columns,
      rows
    } = props.query;

    return { 
      columns,
      rows
     };
  }

  state = {
    columns: this.props.columns,
    rows: this.props.rows
  };

  onGridRowsUpdated = ({ fromRow, toRow, updated }) => {
    this.setState(state => {
      const rows = state.rows.slice();
      for (let i = fromRow; i <= toRow; i++) {
        rows[i] = { ...rows[i], ...updated };
      }
      this.props.update(rows);
      return { rows };
    });
  };

  render() {
    const {
      rows
    } = this.state;
    return (
      <div style={{margin: '20px 0'}}>
        <ReactDataGrid
        columns={this.props.columns}
        rowGetter={i => rows[i]}
        rowsCount={rows.length}
        onGridRowsUpdated={this.onGridRowsUpdated}
        enableCellSelect={true}             
    />

      </div>
    );
  }
}

export default Matrix;