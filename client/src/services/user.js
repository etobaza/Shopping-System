import { useNavigate } from "react-router-dom";
import {useEffect} from "react";

const API_BASE_URL = "http://localhost:8080"; // Replace with your Golang server URL

const register = async (userData) => {
    try {
        const response = await fetch(`${API_BASE_URL}/register`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(userData),
        });

        if (!response.ok) {
            throw new Error("Error registering user");
        }

        const data = await response.json();
        return { success: true, data };
    } catch (error) {
        console.error("Error registering user:", error);
        return { success: false, message: "Registration failed. Please try again." };
    }
};

const login = async (credentials) => {
    try {
        const response = await fetch(`${API_BASE_URL}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(credentials),
        });

        if (!response.ok) {
            throw new Error("Error logging in user");
        }

        const data = await response.json();
        return { success: true, data };
    } catch (error) {
        return { success: false, message: "Incorrect username or password." };
    }
};

export const fetchUserData = async () => {
    const response = await fetch('/user-data');
    const data = await response.json();
    return data;
};

export const fetchCategories = async () => {
    const response = await fetch('/categories');
    const data = await response.json();
    return data;
};

export const handleLogout = async () => {
    try {
        const response = await fetch('/logout', { method: 'POST' });
        if (!response.ok) {
            throw new Error('Failed to logout');
        }
        return true;
    } catch (error) {
        console.error(error);
        return false;
    }
};


export default {
    register,
    login,
};
