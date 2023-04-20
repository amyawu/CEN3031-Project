import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { UploadComponent } from "./upload.component"

describe('UploadComponent', () => {
    it('mounts', () => {
        cy.mount(UploadComponent, {
          imports: [BrowserModule,
            HttpClientModule,
            FormsModule
        ]
        })
      })
      
      it.only('check if you can upload', () => {
        cy.mount(UploadComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
        cy.get('[data-cy=upload_button]').click()
      })

      it.only('check if we can switch to display page from upload', () => {
        cy.mount(UploadComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=switch_display_button_upload]').should('have.attr', 'href', '/display')
    })
    it.only('switches to display page from upload', () => {
        cy.mount(UploadComponent, {
            imports: [BrowserModule,
              HttpClientModule,
              FormsModule
          ]
          })
          cy.get('[data-cy=switch_display_button_upload]').should('have.attr', 'href', '/display')
          cy.get('[data-cy=switch_display_button_upload]').click();
    })
})