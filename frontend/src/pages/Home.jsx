import React from "react"
import Keys from "../components/Keys"
import Signer from "../components/Signer"
import Out from "../components/Out"
import { GetMyPublicKey } from "../../wailsjs/go/main/App"

const Home = () => {
  return (
    <main>
      <Out />
      <div className="grid">
        <Keys />
        <div className="main">
          <Signer />
        </div>
      </div>
    </main>
  )
}

export default Home

export async function loader() {
  const myPublicKey = await GetMyPublicKey()
  return { myPublicKey }
}
