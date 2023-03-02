describe('visiting register spec', () => {
  it('visiting register site', () => {
    cy.visit('http://localhost:4200/register')

  })

  it('fills out the register form and submits it', () => {
    cy.visit('http://localhost:4200/register')
    cy.get('.register_username').type('email@gmail.com')
    cy.get('.register_password').type('emailemail')
    cy.get('.register_button_e2e').click()
  })
  it('check if the user wants to switch to login from register', () => {
    cy.visit('http://localhost:4200/register')
    cy.contains('Login').click()
    cy.url().should('eq','http://localhost:4200/login')
  })
})

describe('visiting login spec', () => {
  it('visiting login site', () => {
    cy.visit('http://localhost:4200/login')
  })

  it('fills out the login form and submits it', () => {
    cy.visit('http://localhost:4200/login')
    cy.get('.login_username').type('email@gmail.com')
    cy.get('.login_password').type('emailemail')
    cy.get('.login_button_e2e').click()
  })

  it('check if the user wants to switch to login from register', () => {
    cy.visit('http://localhost:4200/register')
    cy.contains('Register').click()
    cy.url().should('eq','http://localhost:4200/register')
  })
})

describe('visiting upload spec', () => {
  it('visiting upload site', () => {
    cy.visit('http://localhost:4200/upload')
  })
})