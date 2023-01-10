import React from "react"
import { useRouteError, useNavigate } from "react-router-dom"

const Err = () => {
  const err = useRouteError()
  const navigate = useNavigate()

  return (
    <main>
      <h1>ERROR!</h1>
      <h4 className="error">{JSON.stringify(err)}</h4>
      <p className="link" onClick={() => navigate(-1)}>&larr; Go Back</p>
    </main>
  )
}

export default Err
