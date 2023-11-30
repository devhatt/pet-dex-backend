import "dotenv/config";
import app from "./app";

const port = process.env.API_PORT || "3001";

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});

