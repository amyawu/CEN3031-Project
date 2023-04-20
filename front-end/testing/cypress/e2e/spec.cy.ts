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
})

describe('visiting upload spec', () => {
  it('visiting upload site', () => {
    cy.visit('http://localhost:4200/upload')
  })
  it('check if the user wants to switch to display from upload', () => {
    cy.visit('http://localhost:4200/upload')
    cy.contains('.display_images').click()
    cy.url().should('eq','http://localhost:4200/display')
  })
})

describe('visiting profile spec', () => {
  it('visiting profile site', () => {
    cy.visit('http://localhost:4200/profile')

  })

  it('fills out the profile form and submits it', () => {
    cy.visit('http://localhost:4200/profile')
    cy.get('.profile_name').type('John Doe')
    cy.get('.profile_gender').type('male')
    cy.get('.profile_age').type('21')
    cy.get('.profile_ethnicity').type('white')
    cy.get('.profile_button_e2e').click()
  })
  it('check if the user wants to switch to login from profile', () => {
    cy.visit('http://localhost:4200/profile')
    cy.contains('.login').click()
    cy.url().should('eq','http://localhost:4200/login')
  })
  it('check if the user wants to switch to account from profile', () => {
    cy.visit('http://localhost:4200/profile')
    cy.contains('.display_Info').click()
    cy.url().should('eq','http://localhost:4200/account')
  })
  
})

  describe('visiting account spec', () => {
    it('visiting account site', () => {
      cy.visit('http://localhost:4200/account')
    })
    it('check if the user wants to switch to profile from account', () => {
      cy.visit('http://localhost:4200/account')
      cy.contains('.update').click()
      cy.url().should('eq','http://localhost:4200/profile')
    })
  })

  describe('visiting home spec', () => {
    it('visiting home site', () => {
      cy.visit('http://localhost:4200/home')
    })
    it('visiting home site + testing 1st button', () => {
      cy.visit('http://localhost:4200/home')
      cy.contains('.profile').click()
      cy.url().should('eq','http://localhost:4200/profile')
    })
    it('visiting home site + testing 2nd button', () => {
      cy.visit('http://localhost:4200/home')
      cy.contains('.display').click()
      cy.url().should('eq','http://localhost:4200/display')
    })
    it('visiting home site + + testing 3rd button', () => {
      cy.visit('http://localhost:4200/home')
      cy.contains('.upload').click()
      cy.url().should('eq','http://localhost:4200/upload')
    })
  })


  describe('visiting display spec', () => {
    it('visiting display site', () => {
      cy.visit('http://localhost:4200/display')
    })
    it('check if the user wants to switch to upload from display', () => {
      cy.visit('http://localhost:4200/display')
      cy.contains('.upload').click()
      cy.url().should('eq','http://localhost:4200/upload')
    })
  })