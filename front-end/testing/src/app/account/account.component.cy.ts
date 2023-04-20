import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { AccountComponent } from "./account.component"

describe('Account Page', () => {
    it('mounts', () => {
        cy.mount(AccountComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      

      it.only('check if we can switch to home page from account', () => {
        cy.mount(AccountComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=home_button_account]').should('have.attr', 'href', '/home')
    })
    
    it.only('check if we can switch to profile page from account', () => {
      cy.mount(AccountComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
        cy.get('[data-cy=profile_button_account]').should('have.attr', 'href', '/profile')
  })

  it.only('check if we can switch to image page from account', () => {
    cy.mount(AccountComponent, {
        imports: [BrowserModule,
          HttpClientModule,
          FormsModule
      ]
      })
      cy.get('[data-cy=upload_button_account]').should('have.attr', 'href', '/upload')
})

it.only('check if we can switch to register page from account', () => {
  cy.mount(AccountComponent, {
      imports: [BrowserModule,
        HttpClientModule,
        FormsModule
    ]
    })
    cy.get('[data-cy=register_button_account]').should('have.attr', 'href', '/register')
})

it.only('check if we can switch to login page from account', () => {
  cy.mount(AccountComponent, {
      imports: [BrowserModule,
        HttpClientModule,
        FormsModule
    ]
    })
    cy.get('[data-cy=login_button_account]').should('have.attr', 'href', '/login')
})

it.only('Logged out of account page', () => { // need to fix
  cy.mount(AccountComponent, {
      imports: [BrowserModule,
        HttpClientModule,
        FormsModule
    ]
    })
  cy.get('[data-cy=logout_button_account]').submit()
})
})