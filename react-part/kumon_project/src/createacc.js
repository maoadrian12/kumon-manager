import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'

const CreateAcc = (props) => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const [name, setName] = useState('')
  const [nameError, setNameError] = useState('')

  const navigate = useNavigate()

  const onButtonClick = () => {
    setUsernameError('')
    setPasswordError('')
    setNameError('')

    // Check if the user has entered both fields correctly
    if ('' === username) {
        setUsernameError('Please enter your username')
        return
    }
    if ('' === name) {
        setNameError('Please enter your name')
        return
    }

    if ('' === password) {
        setPasswordError('Please enter a password')
        return
    }

    checkAccountExists(username).then((accountExists) => {
        console.log("result is " + accountExists)
        if (accountExists) window.confirm('User already exists. Please login.')
        else {
          CreateAccount()
        }
      });

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

const CreateAccount = () => {
    fetch('http://localhost:8080/createacc', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "username": username, "pass": password, "name": name }),
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
          navigate('/')
        }
      })
  }



  return (
    <div className={'mainContainer'}>
      <div className={'titleContainer'}>
        <div>Create new account</div>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={username}
          placeholder="Enter your username"
          onChange={(ev) => setUsername(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{usernameError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={name}
          placeholder="Enter your name"
          onChange={(ev) => setName(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{nameError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={password}
          placeholder="Enter your password"
          onChange={(ev) => setPassword(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{passwordError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Create new account'} />
      </div>
      <div>
        <p>Go back <a href="/login">here</a></p>
      </div>
    </div>
  )
}

export default CreateAcc