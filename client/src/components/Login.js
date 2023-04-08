import React, { useState } from "react";
import userService from "../services/user";
import { useNavigate } from "react-router-dom";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const credentials = {
            username,
            password,
        };

        const result = await userService.login(credentials);
        if (result.success) {
            console.log("User logged in:", result.data);
            navigate("/home");
        } else {
            console.error("Error logging in user:", result.message);
        }


    };


    return (
        <div>
            <h1>User Login</h1>
            <form onSubmit={handleSubmit}>
                <label htmlFor="username">Username:</label>
                <input
                    type="text"
                    id="username"
                    name="username"
                    required
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <br/>
                <label htmlFor="password">Password:</label>
                <input
                    type="password"
                    id="password"
                    name="password"
                    required
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <br/>

                <input type="submit" value="Login"/>
            </form>
        </div>
    );
};

export default Login;
