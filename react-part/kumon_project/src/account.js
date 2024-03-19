import React from 'react'
import { useNavigate, useParams } from 'react-router-dom'
const Account = (props) => {
  console.log(props);
  const { loggedIn, username } = props
  const navigate = useNavigate()
  console.log(username)
  console.log(loggedIn)

  const signout = () => {
    if (loggedIn) {
      props.setLoggedIn(false)
      props.setUsername('name')
    } else {
        navigate("/login")
    }
  }
  const addStudent = () => {
    navigate("/addstudent")
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


  const ColoredLine = ({ color }) => (
    <hr
        style={{
            color: color,
            backgroundColor: color,
            height: 10
        }}
    />
);

  return (
    <div className="mainContainer">
      <div className={'titleContainer'}>
        <div>
        <input
            className={'signoutButton'}
            type="button"
            onClick={signout}
            value={loggedIn ? 'Log out' : 'Log out'}
            />
            Welcome {username}!
        <input
            className={'deleteButton'}
            type="button"
            onClick={deleteAcc}
            value={'Delete Account'}
            />
            </div>
        <div className={'smaller'}>Your Students</div>
      </div>
      <hr className="separator"/>
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