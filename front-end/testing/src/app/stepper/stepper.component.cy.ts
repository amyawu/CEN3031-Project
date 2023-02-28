import { StepperComponent } from './stepper.component'

describe('StepperComponent', () => {
  it('mounts', () => {
    cy.mount(StepperComponent)
  })
  it('when the increment button is pressed, the counter is incremented', () => {
    cy.mount(StepperComponent)
    cy.get('[data-cy=increment]').click()
    cy.get('[data-cy=counter]').should('have.text', '1')
  })
  
  it('when the decrement button is pressed, the counter is decremented', () => {
    cy.mount(StepperComponent)
    cy.get('[data-cy=decrement]').click()
    cy.get('[data-cy=counter]').should('have.text', '-1')
  })
})