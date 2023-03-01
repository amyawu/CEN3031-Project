import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { LoginComponent } from "./login.component"

describe('LoginComponent', () => {
    it('mounts', () => {
        cy.mount(LoginComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('login without response', () => {
        cy.mount(LoginComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy=login_form]').submit()
      })

      it.only('fill out login with response', () => {
        cy.mount(LoginComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy="login_email_data"]').type('email@gmail.com')
        cy.get('[data-cy="login_password_data"]').type('emailemail')
    })

    it.only('login with filled out response', () => {
        cy.mount(LoginComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy="login_email_data"]').type('email@gmail.com')
        cy.get('[data-cy="login_password_data"]').type('emailemail')
        cy.get('[data-cy=login_form]').submit()
    })

    it.only('check if we can switch to register page from login', () => {
        cy.mount(LoginComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=register_button_login]').should('have.attr', 'href', '/register')
    })
    it.only('switches to register page from login', () => {
        cy.mount(LoginComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=register_button_login]').should('have.attr', 'href', '/register')
          cy.get('[data-cy=register_button_login]').click();
    })
})