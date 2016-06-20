import React from 'react';
import 'whatwg-fetch';

import CityForecast from './city-forecast';
import CityGlows from './city-glows';

import './index.css';

export default class PolyglotCo extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      backendBuild: null,
      lambdaBuild: null,
      forecasts: []
    };
  }

  componentDidMount() {
    this._fetchForecasts();
    setInterval(() => { this._fetchForecasts() }, 1000);

    this._fetchBackendBuild();
    setInterval(() => { this._fetchBackendBuild() }, 1000);
  }

  render() {
    return (
      <div style={{height:"100vh"}}>
        {this._renderCityGlows()}
        <div className="p3 absolute bottom-0 right-0 z2" style={{mixBlendMode:'screen'}}>
          <p className="h2 m0 caps bold"><a className="color-trans-white text-decoration-none hover-color-white transition-all" href="https://github.com/buildkite/polyglot-co-demo">Github Source →</a></p>
        </div>
        <div className="p3 absolute top-0 left-0 z2">
          {this._renderBackendBuild()}
          {this._renderLambdaBuild()}
        </div>
        {this._renderCityForecasts()}
      </div>
    )
  }

  _renderCityGlows() {
    if (this.state.forecasts.length > 0) {
      return (
        <CityGlows
          temperature1={this.state.forecasts[0].high}
          temperature2={this.state.forecasts[1].high}
          temperature3={this.state.forecasts[2].high}
          temperature4={this.state.forecasts[3].high}/>
      )
    }
  }

  _renderCityForecasts() {
    if (this.state.forecasts.length > 0) {
      return (
        <div className="flex items-center justify-center flex-column z1 relative" style={{height:"100%"}}>
          <div className="flex">
            <CityForecast {...this.state.forecasts[0]} />
            <CityForecast {...this.state.forecasts[1]} />
          </div>
          <div className="flex">
            <CityForecast {...this.state.forecasts[2]} />
            <CityForecast {...this.state.forecasts[3]} />
          </div>
        </div>
      )
    }
  }

  _renderBackendBuild() {
    if (this.state.backendBuild) {
      return (
        <p className="h2 bold m0 caps color-trans-white">
          {`Backend Build #${this.state.backendBuild}`}
        </p>
      )
    }
  }

  _renderLambdaBuild() {
    if (this.state.lambdaBuild) {
      return (
        <p className="h2 bold m0 mt1 caps color-trans-white">
          {`Lambda Build #${this.state.lambdaBuild}`}
        </p>
      )
    }
  }

  _fetchForecasts() {
    fetch("/forecasts")
      .then((response) => {
        return response.json()
      }).then((json) => {
        this.setState({
          forecasts: json["forecasts"],
          lambdaBuild: json["build"]
        })
      }).catch((e) => {
        console.log('parsing failed', e)
      })
  }

  _fetchBackendBuild() {
    fetch("/build")
      .then((response) => {
        return response.json()
      }).then((json) => {
        this.setState({
          backendBuild: json["build"]
        })
      }).catch((e) => {
        console.log('parsing failed', e)
      })
  }
}
