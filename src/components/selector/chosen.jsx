import React from 'react';

class Chosen extends React.Component {
  constructor(args) {
    super(args);
  }
  render() {
    const _onClick = () => {
      this.props.onChosenClicked('Yo');
    };
    return (
      <div onClick={_onClick}
           className={this.props.visible ? "selector__label--visible" : "selector__label--hidden"}
           style={{padding: '10px'}}
           title="Click to Change"
      >
        <span>{this.props.value}</span>
      </div>
    )
  }
}

export default Chosen;