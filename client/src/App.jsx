import React, { useState } from "react"

function App() {


  const [name, setName] = useState("")


  const handleSubmit = async (e) => {
    e.preventDefault()
    const response = await fetch(import.meta.env.VITE_API + '/users', {
      method: 'POST',
      body: JSON.stringify({name})
    })
    const data = await response.json()
    console.log(data)
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>

        <input
          type="name"
          placeholder="Coloca tu nombre"
          onChange={e => setName(e.target.value)}
        />

        <button>Guardar</button>
      </form>
    </div>
  )
}

export default App