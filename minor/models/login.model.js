const mongoose= require('mongoose');

const Login = mongoose.model(
    "Login",
    new mongoose.Schema({
        
        email: String,
        password: String,
        roles: [
            {
                type: mongoose.Schema.Types.ObjectId,
                ref:"Role"
            }
        ]
    })
)

module.exports = Login;