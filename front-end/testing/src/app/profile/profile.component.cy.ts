import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { ProfileComponent } from "./profile.component"

describe('ProfileComponent', () => {
    it('mounts', () => {
        cy.mount(ProfileComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('profile page without response', () => {
        cy.mount(ProfileComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy=profile_form]').submit()
      })

      it.only('fill out profile page with response', () => {
        cy.mount(ProfileComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy="profile_name_data"]').type('John Doe')
        cy.get('[data-cy="profile_gender_data"]').type('male')
        cy.get('[data-cy="profile_age_data"]').type('21')
        cy.get('[data-cy="profile_ethnicity_data"]').type('white')
    })

    it.only('fill out profile page with response and submit', () => {
        cy.mount(ProfileComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy="profile_name_data"]').type('John Doe')
          cy.get('[data-cy="profile_gender_data"]').type('male')
          cy.get('[data-cy="profile_age_data"]').type('21')
          cy.get('[data-cy="profile_ethnicity_data"]').type('white')
        cy.get('[data-cy=profile_form]').submit()
    })
    
    it.only('check if we can switch to home page from login', () => {
      cy.mount(ProfileComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
        cy.get('[data-cy=home_button_login]').should('have.attr', 'href', '/home')
  })
  
  it.only('check if we can switch to profile page from login', () => {
    cy.mount(ProfileComponent, {
        imports: [BrowserModule,
          HttpClientModule,
          FormsModule
      ]
      })
      cy.get('[data-cy=profile_button_login]').should('have.attr', 'href', '/profile')
})

it.only('check if we can switch to image page from login', () => {
  cy.mount(ProfileComponent, {
      imports: [BrowserModule,
        HttpClientModule,
        FormsModule
    ]
    })
    cy.get('[data-cy=upload_button_login]').should('have.attr', 'href', '/upload')
})

it.only('check if we can switch to register page from login', () => {
cy.mount(ProfileComponent, {
    imports: [BrowserModule,
      HttpClientModule,
      FormsModule
  ]
  })
  cy.get('[data-cy=register_button_login]').should('have.attr', 'href', '/register')
})

it.only('check if we can switch to login page from login', () => {
cy.mount(ProfileComponent, {
    imports: [BrowserModule,
      HttpClientModule,
      FormsModule
  ]
  })
  cy.get('[data-cy=login_button_login]').should('have.attr', 'href', '/login')
})
    
})