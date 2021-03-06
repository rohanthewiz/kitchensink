import React from 'react';
import Selector from './selector/selector.jsx';

const QueryLimit = 8; // This must correspond to the limit set on the backend query

class App extends React.Component {
  constructor(args) {
    super(args);
    this.state = {selectValue: 'Nothing chosen',
      // debounce: {timer: null, period: 300, // qe -- queue_empty, ctp -- clear_to_pass
      // results: null, input: null}
    }
  }
  render() {

    const doFetch = (input) => {
      return fetch(`/products?q=${input}`, {credentials: 'include'})
        .then((response) => {
          return response.json();
        }).then((data) => {
          let ret = {options: data};
          if(data.length > 0 && data.length < QueryLimit) ret['complete'] = true;
          return ret;
        });
    };

    return (
      <div id="app">
        <strong>A Simple and easily customized selector for React</strong><br/>
        <span>(data is supplied from SQLite on the backend)</span><br/>
        <span>Click the button below to change the selection.</span>
        <Selector
          loadOptions={doFetch}
          minimumInput={1}
          placeholder="Search products"
          value={this.state.selectValue}
          selectionChanged={this.saveSelectedAsset.bind(this)}
          title="Choose a product"
        />
      </div>)
  }

  saveSelectedAsset(val) {
    console.log("Chosen:");
    console.log(val);
    this.setState({selectValue: val});
  }
}

export default App
