import express, { Request, Response } from "express";
import "dotenv/config";
import { PrismaClient } from "@prisma/client";
const prisma = new PrismaClient();
const app = express();
const port = process.env.PORT || 3001;
app.use(express.json());
app.get("/", (req: Request, res: Response) => {
  res.send("Hello, World!");
});

app.post("/users/create", async (req: Request, res: Response) => {
  try {
    const { name, email } = req.body;
    const user = await prisma.user.create({
      data: {
        name,
        email,
      },
    });
    res.json(user);
  } catch (error) {
    res.status(500).json({ error: "Failed to create user" });
  }
});
app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
