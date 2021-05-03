import './App.css';
import axios from 'axios';
import React,{Component} from 'react';

class App extends Component {
	constructor(props){
		super(props);
		this.state = {
			isLogin : false,
			Name : "",
			ThumbNail: "",	
		}
	}

	componentDidMount() {
		axios.get('/api/index')
			.then(resp => {
				let data = resp["data"];
				console.log(data);
				if (data !== ""){
					
					this.setState({
						isLogin : true,
						Name: data.nickName,
						ThumbNail: data.thumbnailURL,
					})
					
					console.log("IsLoggined")
				} else {
					this.setState({isLogin : false})
				}

			})
	}

	render() {
		return (
			<div className="App">
				<div className="Container">
					<LoginHandler isLogin={this.state.isLogin} Name={this.state.Name} ThumbNail={this.state.ThumbNail}  />
				</div>
			</div>
		)
	}
}
class LoginHandler extends React.Component{

	constructor(props){
		super(props);
	}

	render() {
		return (
			<div className="LoginHandler">
				{this.props.isLogin ? <FriendPage Name={this.props.Name} ThumbNail={this.props.ThumbNail}/> : <IndexPage /> }
			</div>
		)

	}

}
class MyInfo extends React.Component{
	render() {
		return (
			<div className="MyInfo">
				<img className="MyThumbNail ml-2 mr-2" src={this.props.ThumbNail} alt="" /><span className="MyName ml-2">{this.props.Name}</span>
			</div>
		)
	}
}
class FriendPage  extends React.Component{

	render() {
		return (
			<div className="FriendPage col-6 ">
				<div className="ListHead mt-3">
					<h1 className="ListHeadText mt-3 mb-3"><span className="badge bg-warning">KakaoWeb</span>
</h1>				
				</div>
				<div className="ListBody mt-3">
					<div className="UserInfo">
						<MyInfo Name={this.props.Name} ThumbNail={this.props.ThumbNail} />
					</div>
					<div className="FriendList">
					</div>
				</div>
			</div>
		)
	}
}

function IndexPage() {
	return (
		<div className="IndexPage">
			<div className="IndexHeader">
				<h1><b>KaKaoWeb Service</b></h1>
			</div>
			<div className="IndexBody">
				<a href="/api/login">Start with Kakao</a>
			</div>
		</div>
	)
}

export default App;
