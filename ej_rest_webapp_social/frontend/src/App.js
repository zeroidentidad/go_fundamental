import React, { useState } from 'react';
import SignIn from './pages/SignIn';
import { ToastContainer } from "react-toastify";

function App() {

  const [user, setUser]=useState('zero');

  return (
    <div>
      {user? <SignIn/>:<h1>No logueado</h1>}
      <ToastContainer
        position="top-right"
        autoClose={5000}
        hideProgressBar
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnVisibilityChange
        draggable
        pauseOnHover
      />
    </div>
  )

}

export default App;