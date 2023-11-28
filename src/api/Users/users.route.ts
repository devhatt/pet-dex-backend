import { Router } from "express";
import userController from "./users.controller";

const users: Router = Router();

users.get("/", userController.getAllUsers);

export default users;
