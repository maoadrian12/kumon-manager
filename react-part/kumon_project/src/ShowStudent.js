import React , { useState, useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
const ShowStudent = (props) => {

  const [latestMath, setLatestMath] = useState('')
  const [daily, setDaily] = useState('')
  const [latestEnglish, setLatestEnglish] = useState('')
  const [secondSet, setSecondSet] = useState([]);
  const [values, setValues] = useState([]);
  const [filterMath, setFilterMath] = useState('');
  const [filterEng, setFilterEng] = useState('');
  const [minMathWkst, setMinMathWkst] = useState('');
  const [minEngWkst, setMinEngWkst] = useState('');
  const [maxMathWkst, setMaxMathWkst] = useState('');
  const [maxEngWkst, setMaxEngWkst] = useState('');
  const [mathCompletionData, setMathCompletionData] = useState('');
  const [mathGradeData, setMathGradeData] = useState('');
  const [engCompletionData, setEngCompletionData] = useState('');
  const [engGradeData, setEngGradeData] = useState('');

  

  
  const [mathLevels, setMathLevels] = useState([]);
  const [engLevels, setEngLevels] = useState([]);
  useEffect(() => {
    getStudentLevels();
    setMinMathWkst(1);
    setMaxMathWkst(200);
    setMinEngWkst(1);
    setMaxEngWkst(200);
  }, []);

  const getStudentLevels = () => {
    fetch('http://localhost:8080/getlevels', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "name": studentUsername, "parent_username": localStorage.getItem('username')}) })
      .then((r) => r.json())
      .then((parent) => {
        const [readingStr, mathStr] = parent.Message.slice(1, -1).split('] [');
        const readingArr = readingStr.split(' ')
        const mathArr = mathStr.split(' ')
        mathArr.push('All levels')	
        readingArr.push('All levels')	
        setEngLevels(readingArr)
        setMathLevels(mathArr)
        setFilterEng('All levels')
        setFilterMath('All levels')
      })
    }

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

  const changeEnglish = () => {
    console.log(`Level is ${filterEng} and min is ${minEngWkst} and max is ${maxEngWkst}`)
    fetch('http://localhost:8080/stats', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "student_username": studentUsername, "parent_username": localStorage.getItem('username'), "wkst_level": filterEng, "min_wkst_number": minEngWkst, "max_wkst_number": maxEngWkst, "program_name": "READING" })
      })
      .then((r) => r.json())
      .then((parent) => {
        if (parent.Result === true) {
          console.log(parent.Message)
          const values = parent.Message.split(" ");
          setEngCompletionData(values[0]);
          setEngGradeData(values[1]);
        } else {
          window.alert('Error getting student info.');
        }
      })
  }

  const changeMath = () => {
    console.log(`Level is ${filterMath} and min is ${minMathWkst} and max is ${maxMathWkst}`);
    fetch('http://localhost:8080/stats', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ "student_username": studentUsername, "parent_username": localStorage.getItem('username'), "wkst_level": filterMath, "min_wkst_number": minMathWkst, "max_wkst_number": maxMathWkst, "program_name": "MATH" })
      })
      .then((r) => r.json())
      .then((parent) => {
        if (parent.Result === true) {
          console.log(parent.Message)
          const values = parent.Message.split(" ");
          setMathCompletionData(values[0]);
          setMathGradeData(values[1]);
        } else {
          window.alert('Error getting student info.');
        }
      })
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
  useEffect(() => {
    setFilterEng('All levels')
    setFilterMath('All levels')
    changeMath();
    changeEnglish();
  })

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
      <div className={'statistics'}>
        <div className={'math'}>
          <div className={'mathTitle'}>Math worksheets</div>
          <select onChange={(e) => setFilterMath(e.target.value)}>
            <option value=""> -- Level to filter through -- </option>
                  {/* Mapping through each fruit object in our fruits array
                and returning an option element with the appropriate attributes / values.
              */}
              {mathLevels.map((level) => <option value = {level}>{level}</option>)}
          </select>
          <div>Enter worksheets to filter through!</div>
          
          <div className={'wkstFilter'}>
            <input
              value={minMathWkst}
              placeholder="Minimum worksheet"
              onChange={(ev) => setMinMathWkst(Math.min(ev.target.value, maxMathWkst))}
              className={'minBox'}
            />
            <input
              value={maxMathWkst}
              placeholder="Maximum worksheet"
              onChange={(ev) => setMaxMathWkst(Math.max(ev.target.value, minMathWkst))}
              className={'maxBox'}
            />
          </div>
          <div>
            <input
            className={'filterButton'}
            onClick={changeMath}
            type="button"
            value={'Apply filter'}
            />
          </div>
          <div className={'center-text'}>
            <h3>Statistics for {filterMath}</h3>
            <p>Average completion time: {mathCompletionData}</p>
            <p>Average grade: {mathGradeData}</p>
          </div>
        </div>
        <div className={'eng'}>
          <div className={'englishTitle'}>English worksheets</div>
          <select onChange={(e) => setFilterEng(e.target.value)}>
            <option value=""> -- Level to filter through -- </option>
                  {/* Mapping through each fruit object in our fruits array
                and returning an option element with the appropriate attributes / values.
              */}
              {engLevels.map((level) => <option value = {level}>{level}</option>)}
          </select>
          <div>Enter worksheets to filter through!</div>
          <div className={'wkstFilter'}>
            <input
              value={minEngWkst}
              placeholder="Minimum worksheet"
              onChange={(ev) => setMinEngWkst(Math.min(ev.target.value, maxEngWkst))}
              className={'minBox'}
            />
            <input
              value={maxEngWkst}
              placeholder="Maximum worksheet"
              onChange={(ev) => setMaxEngWkst(Math.max(ev.target.value, minEngWkst))}
              className={'maxBox'}
            />
          </div>
          <div>
            <input
            className={'filterButton'}
            onClick={changeEnglish}
            type="button"
            value={'Apply filter'}
            />
          </div>
          <div className={'center-text'}>
            <h3>Statistics for {filterEng}</h3>
            <p>Average completion time: {engCompletionData}</p>
            <p>Average grade: {engGradeData}</p>
          </div>
        </div>
      </div>


      <div className={'buttonContainer bottom_button'}>
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