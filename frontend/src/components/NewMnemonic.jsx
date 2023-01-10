import React from "react"
import { Link, useLoaderData } from "react-router-dom"
import {GenerateMnemonic} from "../../wailsjs/go/main/App"

const NewMnemonic = () => {
  const { mnemonic } = useLoaderData()

  return (
    <div style={{ maxWidth: '80%', margin: '0 auto' }}>
      <h2 className="accent">{mnemonic}</h2>
      <p>
        This is your 12-word mnemonic.
        Make sure to write it down and keep it in a safe place.
        This is all you need to generate your keys and be able to sign documents.
        Go back to the initial screen and use this mnemonic to access.
      </p>
      <br />
      <Link to="/">&larr; Go Back</Link>
    </div>
  )
}

export default NewMnemonic

export async function loader() {
  const mnemonic = await GenerateMnemonic()
  return { mnemonic }
}
