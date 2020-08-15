import React, { useState, useEffect } from 'react';
import SignIn from './pages/SignIn';
import { ToastContainer } from "react-toastify";
import { AuthContext } from './context';
import { isUserLoggedApi } from "./api/auth";

function App() {
  const [user, setUser]=useState(null);
  const [loadUser, setLoadUser]=useState(false);
  const [refreshCheckLogin, setRefreshCheckLogin]=useState(false);

  useEffect(()=>{
    setUser(isUserLoggedApi());
    setRefreshCheckLogin(false);
    setLoadUser(true);
  },
    [refreshCheckLogin]);

  if (!loadUser) return null;

  return (
    <AuthContext.Provider value={user}>
      {user? <h1>Logueado</h1>:<SignIn setRefreshCheckLogin={setRefreshCheckLogin}/>}
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
    </AuthContext.Provider>
  )

}

export default App;