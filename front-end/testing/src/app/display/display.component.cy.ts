import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { DisplayComponent } from "./display.component"

describe('DisplayComponent', () => {
    it('mounts', () => {
        cy.mount(DisplayComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('check if we can switch to upload page from display', () => {
        cy.mount(DisplayComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=switch_upload_button_display]').should('have.attr', 'href', '/upload')
    })
    it.only('switches to login page from profile', () => {
        cy.mount(DisplayComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=switch_upload_button_display]').should('have.attr', 'href', '/upload')
          cy.get('[data-cy=switch_upload_button_display]').click();
    })
})