const express = require('express');
const router = express.Router()
const User = require('../models/user')
const mongoose = require('mongoose')
const db = "mongodb+srv://4n1m4t10n:porfidio@cluster0.2oravn6.mongodb.net/?retryWrites=true&w=majority"

mongoose.connect(db, err => {
    if (err) {
        console.error('Error' + err)
    } else (
        console.log('Connected to mongodb')
    )
})

router.get('/', (req, res) => {
    res.send('From API router')
})

router.post('/register', (req, res) => {
    var user = new User(req.body);
    user.save().then(registeredUser => {
        res.status(200).send(registeredUser)
    })
    .catch(err => {
        res.status(400).send(err);
    });
});

router.post('/login', (req, res) => {
    var userData = req.body

    User.findOne({email: userData.email}, (error, user) => {
        if (error) {
            console.log(error)
        } else {
            if (!user) {
                res.status(401).send('Invalid email')
            } else
            if (user.password !== userData.password) {
                res.status(401).send('Invalid password')
            } else {
                res.status(200).send(user);
            }
        }
    })

})

router.get('/events', (req, res) => {

})

module.exports = router;