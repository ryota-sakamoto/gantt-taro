import axios from "axios"

export default async function (req, res, next) {
    try {
        const result = await axios({
            method: req.method,
            url: `http://api:8080${req.originalUrl}`,
            headers: req.headers,
            data: req.body
        })

        res.writeHead(result.status)
        res.write(JSON.stringify(result.data))
        res.end()
    } catch (e) {
        res.writeHead(e.response.status)
        res.write(JSON.stringify(e.response.data))
        res.end()
    }
}
  