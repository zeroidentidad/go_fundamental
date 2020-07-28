import React, { useState } from 'react';

function App() {

  const [user, setUser]=useState(null);

  return (
    <div>
      {user? (<h1>Logueado</h1>):(<h1>No logueado</h1>)}
    </div>
  )

}

export default App;