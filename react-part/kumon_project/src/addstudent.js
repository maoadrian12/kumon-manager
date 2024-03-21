import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'

const AddStudent = (props) => {
    // When setting the username, also store it in localStorage
  
  
  // In your component, initialize the username state from localStorage
  const parentUsername = localStorage.getItem('username');
  //const [username, setUsername] = useState(localStorage.getItem('username') || '');
  
  // In your JSX, use the onUsernameChange function for the onChange event

  //const {username} = props
  const [childUsername, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [pages, setPages] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const [pagesError, setPagesError] = useState('')
  const [name, setName] = useState('')
  const [nameError, setNameError] = useState('')
  console.log(childUsername)


  const onUsernameChange = (ev) => {
    const newUsername = ev.target.value;
    setUsername(newUsername);
    localStorage.setItem('username', newUsername);
  };

  const navigate = useNavigate()

  const onButtonClick = () => {
    setUsernameError('')
    setPasswordError('')
    setNameError('')

    // Check if the user has entered both fields correctly
    if ('' === childUsername) {
        setUsernameError('Please enter a name')
        return
    }
    if ('' === name) {
        setNameError('Please enter a math starting level')
        return
    }

    if (!/^[A-O]$/.test(name)) {
        setNameError('Please enter a letter between A and O')
        return
    }

    if ('' === password) {
        setPasswordError('Please enter an english starting level')
        return
    }

    if (!/^([A-H][1-2])|(^[I-L])$/.test(password)) {
        setPasswordError('Please enter a letter valid english starting level.')
        return
    }

    if ('' === pages) {
        setPagesError('Please enter how many pages they do per day')
        return
    }

    if (!/^[1-9]$|^10$/.test(pages)) {
        setPagesError('Please enter a number between 1 and 10')
        return
    }

    CreateStudent()

  }

  const checkAccountExists = (childUsername) => {
    return fetch('http://localhost:8080/check', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({childUsername}),
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

const CreateStudent = () => {
    fetch('http://localhost:8080/createchild', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "Name": childUsername, "Parent_username": parentUsername}),
    })
      .then((r) => r.json())
      .then((child) => {
        console.log(child.Result)
        if (child.Result === true) {
            UpdateStudent()
        } else {
            window.alert('Child already exists.')
        }
      })
  }

  const UpdateStudent = () => {
    fetch('http://localhost:8080/updatechild', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "Student_name": childUsername,
            "Parent_name": parentUsername,
            "math_level": name,
            "reading_level": password,
            "pages_per_day": pages,
        }),
    })
    .then((r) => r.json())
    .then((child) => {
        /*fetch('http://localhost:8080/updatechild', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "Student_name": childUsername,
                "Parent_name": parentUsername,
                "level_name": name,
                "program_name": (password.length === 1 ? "MATH" : "READING"),
                "pages_per_day": pages,
            }),
        })*/
        console.log(child.Result)
        if (child.Result === true) {
            window.alert('Child created.')
            navigate('/account')
        } else {
            window.alert('Error creating child')
        }
    })
  }

  return (
    <div className={'mainContainer'}>
      <div className={'titleContainer'}>
        <div>Add a student</div>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
            value={childUsername}
            placeholder="Enter their full name"
            onChange={(ev) => setUsername(ev.target.value)}
            className={'inputBox'}
        />
        <label className="errorLabel">{usernameError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={name}
          placeholder="Enter their math starting level"
          onChange={(ev) => setName(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{nameError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={password}
          placeholder="Enter their reading starting level"
          onChange={(ev) => setPassword(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{passwordError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input
          value={pages}
          placeholder="Enter their pages per day"
          onChange={(ev) => setPages(ev.target.value)}
          className={'inputBox'}
        />
        <label className="errorLabel">{pagesError}</label>
      </div>
      <br />
      <div className={'inputContainer'}>
        <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Add student'} />
      </div>
      <div>
        <p>Go back <a href="/account">here</a></p>
      </div>
    </div>
  )
}
export default AddStudent