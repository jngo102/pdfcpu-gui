import React, { Component } from 'react'
import { DecryptPDF, EncryptPDF } from "../wailsjs/go/main/App"
import { Button, Grid, TextField, Typography } from '@mui/material'

class App extends Component {
    render() {
        return (
            <div className="App">
                <Grid container spacing={2}>
                    <Grid item xs={12}>
                        <TextField id="pdf-path-input" label="Enter PDF Path" variant="outlined" />
                    </Grid>
                    <Grid item xs={6}>
                        <Button variant="contained" onClick={this.encryptPDF}>Encrypt</Button>
                    </Grid>
                    <Grid item xs={6}>
                        <Button variant="contained" onClick={this.decryptPDF}>Decrypt</Button>
                    </Grid>
                    <Grid item xs={12}>
                        <Typography id="log-text" variant="body1" />
                    </Grid>
                </Grid>
            </div>
        );
    }

    decryptPDF() {
        const pdfInput = document.getElementById("pdf-path-input") as HTMLInputElement
        const pdfPath = pdfInput.value
        console.log("Decrypting PDF: " + pdfPath)
        DecryptPDF(pdfPath).then((res: any) => {
            const resText = res as string
            const logText = document.getElementById("log-text");
            if (logText != null) {
                logText.innerHTML = resText
            }
        })
    }

    encryptPDF() {
        const pdfInput = document.getElementById("pdf-path-input") as HTMLInputElement
        const pdfPath = pdfInput.value
        console.log("Encrypting PDF: " + pdfPath)
        EncryptPDF(pdfPath).then((res: any) => {
            const resText = res as string
            const logText = document.getElementById("log-text");
            if (logText != null) {
                logText.innerHTML = resText
            }
        })
    }
}

export default App
