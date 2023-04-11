import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Register from "./components/Register";
import Login from "./components/Login";
import NoPage from "./components/NoPage";
import Shop from "./components/Shop";
import AdminPanel from "./components/AdminPanel";
import CategoryList from "./components/CategoryList";
import PostItem from "./components/PostItem";
import Cart from "./components/Cart";

export default function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="register" element={<Register />} />
                <Route path="login" element={<Login />} />
                <Route path="shop" element={<Shop />} />
                <Route path="shop/c/:category" element={<CategoryList />} />
                <Route path="shop/items/new" element={<PostItem />} />
                <Route path="/cart" element={<Cart />} />
                <Route path="admin-panel" element={<AdminPanel />} />
                <Route path="*" element={<NoPage />} />
            </Routes>
        </BrowserRouter>
    );
}

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<App />);