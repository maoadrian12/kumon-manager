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