import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HomeComponent } from "./home.component"

describe('HomeComponent', () => {
    it('mounts', () => {
        cy.mount(HomeComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })

      it.only('checks first button of home page', () => {
        cy.mount(HomeComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        
          cy.get('[data-cy=home_profile_button_home]').should('have.attr', 'href', '/profile')
    })
    it.only('switches to said button from home page', () => {
      cy.mount(HomeComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
        cy.get('[data-cy=home_profile_button_home]').should('have.attr', 'href', '/profile')
        cy.get('[data-cy=home_profile_button_home]').click();
  })
    it.only('checks second button of home page', () => {
        cy.mount(HomeComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          
          cy.get('[data-cy=home_recents_button_home]').should('have.attr', 'href', '/display')
    })
    it.only('switches to said button from home page', () => {
      cy.mount(HomeComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
        cy.get('[data-cy=home_recents_button_home]').should('have.attr', 'href', '/display')
        cy.get('[data-cy=home_recents_button_home]').click();
  })
    it.only('checks third button of home page', () => {
        cy.mount(HomeComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=home_upload_button_home]').should('have.attr', 'href', '/upload')
    })
    it.only('switches to said button from home page', () => {
        cy.mount(HomeComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=home_upload_button_home]').should('have.attr', 'href', '/upload')
          cy.get('[data-cy=home_upload_button_home]').click();
    })
})