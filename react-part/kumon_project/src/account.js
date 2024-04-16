import React , { useState, useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
const Account = (props) => {

  const [children, setChildren] = useState([]);
  const [selectedStudent, setSelectedStudent] = useState('');

  const getStudents = () => {
    return fetch('http://localhost:8080/students', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "Parent_username": localStorage.getItem('username')}),
    })
      .then((r) => r.json())
      .then((parent) => {
        console.log(parent.Result)
        if (parent.Result === false) {
          window.alert('Error getting students')
        } else {
          setChildren(parent)
          console.log("children is", children)
        }
      })
  }
  
  useEffect(() => {
    getStudents();
  }, []);

  const [childUsername, setChildUsername] = useState('')
  const [usernameError, setUsernameError] = useState('')
  const { loggedIn} = props
  const username = localStorage.getItem('username')
  const navigate = useNavigate()

  const signout = () => {
    if (loggedIn) {
      props.setLoggedIn(false)
      props.setUsername('name')
      localStorage.setItem('username', '');
    } else {
        localStorage.setItem('username', '');
        navigate("/login")
    }
  }
  const addStudent = () => {
    navigate("/addstudent")
  }

  const checkStudent = () => {

    if (selectedStudent === '') {
      window.alert('Please select a student');
      return;
    }
    navigate("/student/" + selectedStudent)
  }



  const deleteAcc = () => {
    if (window.prompt('Please enter your username to confirm') === username) {
        console.log('deleting ' + username)
        fetch('http://localhost:8080/delete', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "username": username }),
    })
    .then((r) => r.json())
    .then((parent) => {
        console.log(parent.Result)
        if (parent.Result === true) {
            window.alert('Account deleted')
            props.setLoggedIn(false)
            props.setUsername('name')
            localStorage.setItem('username', '');
            navigate('/login')
        } else {
            window.alert('Error deleting account')
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
    }
  }

  return (
    <div className="mainContainer">
      <div classname={'titleContainer'}>
        <div className={'titleContainer'}>
          <div className={'smaller'}>Your Students</div>
        </div>
        <div className={'signoutButton2'}>
          <input
            className={'signoutButton2'}
            type="button"
            onClick={signout}
            value={'Go back'}
          />
        </div>
        <div className={'studentName'}>
          <h1>
          Welcome {username}!
          </h1>
        </div>
          <div className={'deleteButton2'}>
            <input
              className={'deleteButton2'}
              type="button"
              onClick={deleteAcc}
              value={'Delete Account'}
            />
          </div>
      </div>
      <hr className="separator"/>
      <div className={'studentContainer'}>
        <div className={'buttonContainer'}>
        <select onChange={(e) => setSelectedStudent(e.target.value)}> 
            <option value=""> -- Select a child -- </option>
                  {/* Mapping through each fruit object in our fruits array
                and returning an option element with the appropriate attributes / values.
              */}
              {children.map((child) => <option value = {child.Name}>{child.Name}</option>)}
          </select>
          <label className="errorLabel" style={{color: 'red'}}>{usernameError}</label>
          <br></br>
          <input
            className={'inputButton'}
            type="button"
            onClick={checkStudent}
            value={'Go to student'}
          />
        </div>
      </div>
      <div className={'buttonContainer'}>
        <input
          className={'inputButton'}
          type="button"
          onClick={addStudent}
          value={loggedIn ? 'Log out' : 'Add a student'}
        />
        {loggedIn ? <div>Your username is {username}</div> : <div />}
      </div>
    </div>
  )
}

export default Account