import ReactDataGrid from "react-data-grid";
import React, { Component } from 'react';

class Matrix extends Component {
  state = {
    prividersCount: this.props.prividersCount,
    consumersCount: this.props.consumersCount,
    columns: this.props.columns,
    rows: this.props.rows
  };

  onGridRowsUpdated = ({ fromRow, toRow, updated }) => {
    this.setState(state => {
      const rows = state.rows.slice();
      for (let i = fromRow; i <= toRow; i++) {
        rows[i] = { ...rows[i], ...updated };
      }
      return { rows };
    });
  };

  render() {
    const {
      columns,
      rows
    } = this.state;
    return (
      <div style={{margin: '20px 0'}}>
        <ReactDataGrid
        columns={columns}
        rowGetter={i => rows[i]}
        rowsCount={rows.length}
        onGridRowsUpdated={this.onGridRowsUpdated}
        enableCellSelect={true}
        minHeight={35+35*rows.length+"px"}
    />

      </div>
    );
  }
}

export default Matrix;