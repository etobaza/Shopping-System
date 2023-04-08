// src/components/Register.js

import React, { useState } from "react";
import userService from "../services/user";

const Register = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email, setEmail] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [address, setAddress] = useState("");
    const [phone, setPhone] = useState("");
    const [usertype, setUsertype] = useState("customer");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const userData = {
            username,
            password,
            email,
            firstname: firstName,
            lastname: lastName,
            address,
            phone,
            usertype,
        };

        try {
            const newUser = await userService.register(userData);
            console.log("User registered:", newUser);
        } catch (error) {
            console.error("Error registering user:", error);
        }
    };

    return (
        <div>
            <h1>User Registration</h1>
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
                <br />

                <label htmlFor="password">Password:</label>
                <input
                    type="password"
                    id="password"
                    name="password"
                    required
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <br />

                <label htmlFor="email">Email:</label>
                <input
                    type="email"
                    id="email"
                    name="email"
                    required
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                <br />

                <label htmlFor="usertype">User Type:</label>
                <select
                    id="usertype"
                    name="usertype"
                    required
                    value={usertype}
                    onChange={(e) => setUsertype(e.target.value)}
                >
                    <option value="customer">Customer</option>
                    <option value="seller">Seller</option>
                </select>
                <br />

                <label htmlFor="firstname">First Name:</label>
                <input
                    type="text"
                    id="firstname"
                    name="firstname"
                    required
                    value={firstName}
                    onChange={(e) => setFirstName(e.target.value)}
                />
                <br />

                <label htmlFor="lastname">Last Name:</label>
                <input
                    type="text"
                    id="lastname"
                    name="lastname"
                    required
                    value={lastName}
                    onChange={(e) => setLastName(e.target.value)}
                />
                <br />

                <label htmlFor="address">Address:</label>
                <input
                    type="text"
                    id="address"
                    name="address"
                    required
                    value={address}
                    onChange={(e) => setAddress(e.target.value)}
                />
                <br />

                <label htmlFor="phone">Phone:</label>
                <input
                    type="text"
                    id="phone"
                    name="phone"
                    required
                    value={phone}
                    onChange={(e) => setPhone(e.target.value)}
                />
                <br />

                <input type="submit" value="Register" />
            </form>
        </div>
    );
};

export default Register;
