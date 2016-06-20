import React from 'react';
import ReactCSSTransitionGroup from 'react-addons-css-transition-group';

export default class CityGlows extends React.Component {
  static propTypes = {
    temperature1: React.PropTypes.number.isRequired,
    temperature2: React.PropTypes.number.isRequired,
    temperature3: React.PropTypes.number.isRequired,
    temperature4: React.PropTypes.number.isRequired
  };

  render() {
    return (
      <svg viewBox="0 0 100 100" style={{width:"100%",height:"100%"}} className="absolute top-0 left-0 z0">
        <defs>
          {this._radialGradient("cold", "#8DE8FB")}
          {this._radialGradient("cool", "#14C0EB")}
          {this._radialGradient("warm", "#EBBD2B")}
          {this._radialGradient("hot",  "#FF8321")}
        </defs>
        <g>
          {this._circle(0,   0,   this.props.temperature1)}
          {this._circle(100, 0,   this.props.temperature2)}
          {this._circle(0,   100, this.props.temperature3)}
          {this._circle(100, 100, this.props.temperature4)}
        </g>
      </svg>
    )
  }

  _radialGradient(temperatureName, color) {
    return (
      <radialGradient id={`city-glows-gradient-${temperatureName}`}>
        <stop offset="0" stopColor={color}/>
        <stop offset="100" stopColor="#black"/>
      </radialGradient>
    )
  }

  _circle(x, y, temperature) {
    const gradientId = this._radialGradientId(temperature);
    return (
      <ReactCSSTransitionGroup component="g" transitionName="city-glows-circle" transitionEnterTimeout={3500} transitionLeaveTimeout={3500}>
        <circle
          cx={x} cy={y} fill={`url(${gradientId})`} key={`${x}${y}${gradientId}`}
          r={100} opacity={0.7} style={{mixBlendMode:'screen',willChange:'opacity'}} />
      </ReactCSSTransitionGroup>
    )
  }

  _radialGradientId(temperature) {
    if (temperature < 15) {
      return "#city-glows-gradient-cold";
    } else if (temperature < 25) {
      return "#city-glows-gradient-cool";
    } else if (temperature < 25) {
      return "#city-glows-gradient-warm";
    } else {
      return "#city-glows-gradient-hot";
    }
  }
}
