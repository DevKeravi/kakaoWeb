import './App.css';
import axios from 'axios';
import React,{Component} from 'react';
import Login from './containers/Login';


export default class App extends Component {
	render() {
		return (
		<div className="AppContainer">
			<Login></Login>
			</div>
		)
	}

}
