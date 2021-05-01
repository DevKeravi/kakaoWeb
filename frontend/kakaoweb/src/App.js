import './App.css';

function App() {
  return (
    <div className="App">
		<div className="Container">
			<LoginLink />
		</div>
   </div>
  );
}

function LoginLink() {
	return (
	<div className="LoginLink">
		Logged in with <a href="/api/login">kakao</a>
		</div>
	)
}

export default App;
