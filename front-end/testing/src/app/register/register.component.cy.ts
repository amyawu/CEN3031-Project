import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { RegisterComponent } from "./register.component"

describe('RegisterComponent', () => {
    it('mounts', () => {
        cy.mount(RegisterComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('register without response', () => {
        cy.mount(RegisterComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy=register_form]').submit()
      })

      it.only('fill out login with response', () => {
        cy.mount(RegisterComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy="register_email_data"]').type('email@gmail.com')
        cy.get('[data-cy="register_password_data"]').type('emailemail')
    })

    it.only('fill out login with response', () => {
        cy.mount(RegisterComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy="register_email_data"]').type('email@gmail.com')
        cy.get('[data-cy="register_password_data"]').type('emailemail')
        cy.get('[data-cy=register_form]').submit()
    })

    it.only('check if we can switch to register page from login', () => {
        cy.mount(RegisterComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=login_button_register]').should('have.attr', 'href', '/login')
    })
    it.only('switches to register page from login', () => {
        cy.mount(RegisterComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=login_button_register]').should('have.attr', 'href', '/login')
          cy.get('[data-cy=login_button_register]').click();
    })
})