import React,{Component} from 'react';
import store from '../store' ;

export default class Login extends Component {
	state = {test:1}

	render() {
		return (
			<div className="LoginPage">
				<div className="IndexHeader">
					<h1><b>KakaoWeb Service</b></h1>
				</div>
				<div className="IndexBody">
					<a href="/api/login">Start with kakao</a>
				</div>
			</div>
		)
	}

}
