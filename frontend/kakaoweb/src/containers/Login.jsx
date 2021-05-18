import React,{ Component } from 'react';
import Login from '../components/Login';
import { connect } from 'react-redux';

function mapDispatchToProps(dispatch) {
	return {
		onClick: function(size){
			dispatch({type:'INCREMENT',size:size});
		}
	}
}

export default connect(null,mapDispatchToProps)(Login);
