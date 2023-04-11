import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HomeCmponent } from "./home.component"

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

    it.only('check if we can switch to login page from profile', () => {
        cy.mount(ProfileComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=login_button_register]').should('have.attr', 'href', '/login')
    })
    it.only('switches to login page from profile', () => {
        cy.mount(ProfileComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=login_button_register]').should('have.attr', 'href', '/login')
          cy.get('[data-cy=login_button_register]').click();
    })
})