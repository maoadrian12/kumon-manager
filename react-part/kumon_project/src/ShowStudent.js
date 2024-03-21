import React , { useState, useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
const ShowStudent = (props) => {

  const [latestMath, setLatestMath] = useState('')
  const [daily, setDaily] = useState('')
  const [latestEnglish, setLatestEnglish] = useState('')
  const [secondSet, setSecondSet] = useState([]);
  const [values, setValues] = useState([]);

  useEffect(() => {
    if (latestMath) {
      const values = latestMath.split(" ");
      setSecondSet(values.slice(0, values.length / 2));
      setFirstSet(values.slice(values.length / 2, values.length));
    }
  }, [latestMath]);
  const [firstSet, setFirstSet] = useState([]);
  const { studentUsername } = useParams();
  useEffect(() => {
    getStudentWkst();
  }, []);
  useEffect(() => {
    perDay();
  }, []);

  const getStudentWkst = () => {
    fetch('http://localhost:8080/getinfo', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "name": studentUsername, "parent_username": localStorage.getItem('username')}) })
    .then((r) => r.json())
    .then((parent) => {
        console.log(parent.Result)
        if (parent.Result === false) {
          window.alert('Error getting student info.')
          console.log("error 1")
        } else {
          console.log("success1");
          setLatestEnglish(parent.Message);
          setLatestMath(parent.Message);
          console.log(parent.Message);
        }
    })
    if (latestMath) {
      const values = latestMath.split(" ");
      setSecondSet(values.slice(0, values.length / 2));
      setFirstSet(values.slice(values.length / 2, values.length));
    }
  }

  var perDay = () => {
    fetch('http://localhost:8080/getpages', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "name": studentUsername, "parent_username": localStorage.getItem('username')})
    })
    .then((r) => r.json())
    .then((parent) => {
      if (parent.Result === true) {
        setDaily(parent.Message);
      } else {
        window.alert('Error getting student info.');
        navigate('/account');
      }
    })
  }
  const navigate = useNavigate()

  const signout = () => {
        navigate("/account")
  }

  const complete = (params) => {
    fetch('http://localhost:8080/complete', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "student_username": params.studentUsername, "parent_username": params.parentUsername, "wkst_number": params.worksheetNumber, "wkst_level": params.levelName, "program_name": params.programName, "completion_time": params.time, "grade": params.score}) })
    .then((r) => r.json())
    .then((parent) => {
        if (parent.Result === true) {
        } else {
        }
    })
  }

  const addWorksheets = () => {
    console.log("perday is", daily)
    var time = window.prompt('How many minutes did your student work today?');
    while (isNaN(time) || time <= 0 || time > 100) {
      time = window.prompt('Invalid input. Please enter a number between 1 and 100.');
    }
    time = time / daily
    for (let i = 0; i < daily; i++) {
      var score = window.prompt('Enter the score for the math worksheet number '+ + (Number(secondSet[2]) + i + 1));
      while (isNaN(score) || score < 0 || score > 100) {
        score = window.prompt('Invalid input. Please enter a number between 0 and 100.');
      }
      complete({
        studentUsername: studentUsername,
        parentUsername: localStorage.getItem('username'),
        worksheetNumber: String(Number(secondSet[2]) + i + 1),
        levelName: secondSet[3],
        programName: 'READING',
        time: time,
        score: score
      });}
    for (let i = 0; i < daily; i++) {
      var score = window.prompt('Enter the score for the english worksheet number ' + (Number(firstSet[2]) + i + 1));
      while (isNaN(score) || score < 0 || score > 100) {
        score = window.prompt('Invalid input. Please enter a number between 0 and 100');
      }
      complete({
        studentUsername: studentUsername,
        parentUsername: localStorage.getItem('username'),
        worksheetNumber: String(Number(firstSet[2]) + i + 1),
        levelName: firstSet[3],
        programName: 'MATH',
        time: String(time),
        score: String(score)
      });}
      getStudentWkst();
  }



  const deleteAcc = () => {
    if (window.prompt('Please enter the student\'s name to confirm') === studentUsername) {
        console.log('deleting ' + studentUsername)
        fetch('http://localhost:8080/deletstudent', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "name": studentUsername, "parent_username": localStorage.getItem('username')}) })
    .then((r) => r.json())
    .then((parent) => {
        console.log(parent.Result)
        if (parent.Result === true) {
            window.alert('Student deleted')
            navigate('/account')
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
      <div className={'titleContainer'}>
        <div>
        <input
            className={'signoutButton'}
            type="button"
            onClick={signout}
            value={'Go back'}
            />
            Student {studentUsername}
        <input
            className={'deleteButton'}
            type="button"
            onClick={deleteAcc}
            value={'Delete Student'}
            />
            </div>
      </div>
      <hr className="separator"/>
      <div className={'studentContainer'}>
        Your student's latest work:
      </div>
      <div className={'workContainer'}>
        English: {secondSet[3] + " " + secondSet[2]}
      </div>
      <div className={'workContainer'}>
        Math: {firstSet[3] + " " + firstSet[2]}
      </div>
      <div className={'buttonContainer'}>
        <input
          className={'inputButton'}
          onClick={addWorksheets}
          type="button"
          value={'Complete worksheets'}
        />
      </div>
    </div>
  )
}

export default ShowStudent