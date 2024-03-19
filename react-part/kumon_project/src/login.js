import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'

const Login = (props) => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')

  const navigate = useNavigate()

  const onButtonClick = () => {
    setUsernameError('')
    setPasswordError('')

    // Check if the user has entered both fields correctly
    if ('' === username) {
        setUsernameError('Please enter your username')
        return
    }

    if ('' === password) {
        setPasswordError('Please enter a password')
        return
    }

    checkAccountExists(username).then((accountExists) => {
        console.log("result is " + accountExists)
        if (accountExists) logIn()
        else {
          window.confirm('User does not exist.')
        }
      });
  }

  const logIn = () => {
    fetch('http://localhost:8080/parent', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "username": username, "pass": password }),
    })
      .then((r) => r.json())
      .then((parent) => {
        console.log(parent.Result)
        if (parent.Result === '') {
            window.alert('Wrong password')
          
        } else {
          //localStorage.setItem('user', JSON.stringify({ username, token: parent.token }))
          //props.setLoggedIn(true)
          //props.setUsername(username)
          props.setUsername(username)
          navigate('/account')
        }
      })
  }

  const checkAccountExists = (username) => {
    return fetch('http://localhost:8080/check', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({username}),
    })
      .then((r) => r.json())
      .then((parent) => {
        if (parent.Result === false) {
          console.log('Parent does not exist');
          return false;
        } else {
          console.log('Parent exists');
          return true;
        }
      })
}

  return (
    <div className={'mainContainer'}>
      <div className={'titleContainer'}>
        <div>Login</div>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={username}
          placeholder="Enter your username here"
          onChange={(ev) => setUsername(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{usernameError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={password}
          placeholder="Enter your password here"
          onChange={(ev) => setPassword(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{passwordError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Log in'} />
      </div>
      <div>
        <p>Create a new account <a href="/createacc">here</a></p>
      </div>
    </div>
  )
}

export default Login