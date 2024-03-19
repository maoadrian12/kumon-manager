import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Home from './home'
import Login from './login'
import CreateAcc from './createacc'
import Account from './account'
import './App.css'
import { useEffect, useState } from 'react'

function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  const [username, setUsername] = useState('')

  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn} />} />
          <Route path="/login" element={<Login setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
          <Route path="/createacc" element={<CreateAcc setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
          <Route path="/account" element={<Account username={username} setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
          <Route path="/addstudent" element={<Account username={username} setLoggedIn={setLoggedIn} setUsername={setUsername} />} />
        </Routes>
      </BrowserRouter>
    </div>
  )
}
export default App