import React from "react"
import { useRouteError } from "react-router-dom"

const Err = () => {
  const err = useRouteError()
  return (
    <main>
      <h1>ERROR!</h1>
      <h4 className="error">{JSON.stringify(err)}</h4>
    </main>
  )
}

export default Err
