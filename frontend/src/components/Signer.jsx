import React, { useState } from "react"
import toast from "react-simple-toasts"
import { SelectFile, SignFile, CopyToClipboard } from "../../wailsjs/go/main/App"

const Signer = () => {
  const [path, setPath] = useState("")
  const [signature, setSignature] = useState("")
  const [err, setErr] = useState("")

  const selectFile = async () => {
    try {
      const p = await SelectFile() 
      setPath(p)
    } catch (err) {
      setErr(err)
    } finally {
      setSignature("")
    }
  }

  const signFile = async () => {
    setErr("")
    if (!path) {
      setErr("Please choose a file first")
      setSignature("")
      return
    }
    try {
      const s = await SignFile(path)
      setSignature(s)
    } catch (err) {
      setErr(err)
      setSignature("")
    }
  }

  const copySignatureToClipboard = async () => {
    try {
      await CopyToClipboard(signature)
      toast("Copied to clipboard!")
    } catch (err) {
      setErr(err)
    }
  }

  return (
    <div>
      <section>
        <h2>Select File To Sign</h2>
        <div className="file-chooser">
          <button className="outline" onClick={selectFile}>Open</button>
          <div>{path}</div>
        </div>
        <div className="align-right">
          <button className="fill" onClick={signFile}>Sign File</button>
        </div>
      </section>
      
      <section>
        <h2>Signature</h2>
        <div className="signature">
        {signature}
        </div>
        {signature ? (
          <div className="align-right">
            <button className="fill" onClick={copySignatureToClipboard}>Copy Signature</button>
          </div>
        ) : null}
      </section>

      {err ? <p className="error">{err}</p> : null}
    </div>
  )
}

export default Signer
