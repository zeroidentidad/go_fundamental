import React, { useState } from 'react';
import SignIn from './pages/SignIn';

function App() {

  const [user, setUser]=useState('zero');

  return (
    <div>
      {user? <SignIn/>:<h1>No logueado</h1>}
    </div>
  )

}

export default App;