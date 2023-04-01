import React, {useState} from "react";
import userService from "../services/user";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const credentials = {
            username,
            password,
        };

        try {
            const loggedInUser = await userService.login(credentials);
            console.log("User logged in:", loggedInUser);
        } catch (error) {
            console.error("Error logging in user:", error);
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