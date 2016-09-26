import React from 'react';
import ReactDOM from 'react-dom';
import _ from 'lodash';
import Chosen from './chosen.jsx';
import Dropdown from './dropdown.jsx';

// props
// - loadOptions (method) - for async to be implemented
// - selectionChanged (method) - call when a new value has been selected
// - minimumInput
// - placeholder, title
// state
// - value - the chosen value
// - dropdown_visible - bool
// - items - the items found
// Caching
// If pattern starts with a cached pattern then use cached results

class Selector extends React.Component {
  constructor(args) {
    super(args);
    this.state = {
      inputVisible: false,
      dropdown_visible: false,
      searching: false,
      cache: {}, // {'john br': [result, ...]}
      value: 'Nothing Chosen',
      debounce: {timer: null, period: 300, // qe -- queue_empty, ctp -- clear_to_pass
        results: null, input: null}
    };
  }

  render() {
    // Methods placed here are part of the render closure and so can be passed straight to children

    // Callback for Drop down clicked
    const setSelected = (val) => {
      this.setState({value: val, inputVisible: false, dropdown_visible: false});
      console.log("Selection: ", val);
      // don't call this yet //this.props.selectionChanged(val); // todo
    };
    // Callback for label clicked
    const hideChosen = (arg) => {
      console.log("Clicking label...event:"); console.log(arg);
      //this.refs.selector_input.focus();
      this.setState({inputVisible: true});
      setTimeout(function() { // React state updates are not immediate so go asynch to buy time
        document.querySelector(".selector__input").focus();
        //no workie: ReactDOM.findDOMNode(this.refs.selector_input).focus();
      }, 0);

      return false;
    };

    return (
      <div className="selector">
        <input type="text" autoFocus
               className={this.state.inputVisible ?
                 "selector__input selector__input--visible" :
                 "selector__input selector__input--hidden"}
               title={this.props.title}
               placeholder={this.props.placeholder}
               onChange={this.inputChanged.bind(this)}
               onKeyDown={this.inputKeyDown.bind(this)}
               ref="selector_input"
        />
        <Chosen value={this.state.value}
               visible={!this.state.inputVisible}
               onChosenClicked={hideChosen}
        />
        <Dropdown className="selector__dropdown"
           items={this.state.items}
           searching={this.state.searching}
           visible={this.state.dropdown_visible}
           setSelected={setSelected}
           ref="selector_dropdown"
        />
      </div>
    )
  }

  // EVENT HANDLERS

  inputKeyDown(evt) {
    let dropdown = this.refs.selector_dropdown;
    if (evt.keyCode == 13) { // Enter
      if (this.state.items.length < 1) return false;
      let selIdx = dropdown.selectedIndex();
      if (selIdx >= 0) {
        console.log("Our selection is: ", this.state.items[selIdx]);
        this.setState({value: this.state.items[selIdx].label,
          inputVisible: false
        });
        // Not yet: this.props.selectionChanged(this.state.items[selIdx]); // TODO
      } else {
        console.log("Nothing selected");
      }
      evt.preventDefault();
      this.setState({dropdown_visible: false});
      return false;
    }
    if (evt.keyCode == 40) { // Down Arrow
      dropdown.incrementSelectedIndex();
    }
    if (evt.keyCode == 38) { // Up Arrow
      dropdown.decrementSelectedIndex();
    }
    return true;
  }

  inputChanged(evt) {
    var val = evt.target.value;
    if (this.isBlank(val)) {
      this.setDropdownVisible(false); return true;
    }
    if (val.length < this.props.minimumInput) return true;
    this.setState({value: val}); // dbl chk if we need this
    // Filter the list or Make Async call
    let key = this.cached(val);
    if(key) {
      console.log("cache_key is: " + key + ", cached items: " + this.state.cache[key].length);
      this.filterItems(val, key);
      this.setDropdownVisible(true);
    } else {
      setTimeout(() => {
        this.setState({searching: true});
        this.props.loadOptions(val).then((res) => {
          console.log("Async results:"); console.log(res);
          if(res.options) {
            this.setState({items: res.options, searching: false});
            this.setDropdownVisible(true); // to do: visible based on items length or "no result"
            // Cache if less than limit values returned
            if(res.complete) {
              if(!this.cached(val)) {
                // `[val]` see https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Object_initializer#Computed_property_names
                let new_cache = _.assign(this.state.cache, {[val]: res.options});
                this.setState({cache: new_cache});
              }
            }
          }
        });
      }, 0);
    }
    return true;
  }

  // UTILS

  cached(pattern) {
    let keys = _.keys(this.state.cache);
    for(let i=0; i < keys.length; i++) {
      if(_.startsWith(pattern, keys[i])) {
        console.log("Matching cache key found: '" + keys[i] + "'");
        return keys[i];
      }
    }
    return false;
  }

  setDropdownVisible(show) { // show --> bool
    if(typeof(show) === "undefined") show = true;
    if(show && !this.state.dropdown_visible) this.setState({dropdown_visible: true});
    if(!show && this.state.dropdown_visible) this.setState({dropdown_visible: false});
  }

  // Local filtering of Items
  filterItems(pattern, key) {
    var regstr = _.filter(_.escapeRegExp(pattern).split(' '), (token) => {
      return token.length >= 1;
    }).join('.*\\s');
    let rgx = new RegExp(regstr, "i");
    console.log("regstr: " + regstr);

    let filtered = _.filter(this.state.cache[key], (item) => {
      return item.label.search(rgx) >= 0;
    });
    this.setState({items: filtered});
  }

  isBlank(str) {
    return _.trim(str) === '';
  }
}

export default Selector;
