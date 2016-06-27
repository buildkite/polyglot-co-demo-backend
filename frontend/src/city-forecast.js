import React from 'react';
import ReactCSSTransitionGroup from 'react-addons-css-transition-group';

export default class CityForecast extends React.Component {
  static propTypes = {
    name: React.PropTypes.string.isRequired,
    high: React.PropTypes.number.isRequired,
    summary: React.PropTypes.string.isRequired
  };

  render() {
    return (
      <div className="p2 relative" style={{width:'15rem'}}>
        <h3 className="m0 h2 color-white" style={{letterSpacing:'0.05em',marginBottom:"-0.3em"}}>{this.props.name}</h3>
        <ReactCSSTransitionGroup transitionName="city-forecast-temperature" transitionEnterTimeout={500} transitionLeaveTimeout={500}>
          <p className="m0 bold absolute" style={{fontSize:'8rem',letterSpacing:'-0.015em',opacity:0.999,willChange:'opacity'}} key={Math.floor(this.props.high)}>
            {`${Math.floor(this.props.high)}°`}
          </p>
        </ReactCSSTransitionGroup>
        <p className="m0 h4 truncate color-white" style={{marginTop:'7.8rem',paddingBottom:".25em"}}>{this.props.summary}</p>
      </div>
    )
  }
}
