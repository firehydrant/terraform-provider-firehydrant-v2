resource "firehydrant_signals_api_escalation_policy_entity" "my_signals_api_escalationpolicyentity" {
  default     = true
  description = "...my_description..."
  handoff_step_input = {
    target_id   = "...my_target_id..."
    target_type = "...my_target_type..."
  }
  name = "...my_name..."
  prioritized_settings = {
    high = {
      handoff_step = {
        target_id   = "...my_target_id..."
        target_type = "...my_target_type..."
      }
      repetitions = 0
    }
    low = {
      handoff_step = {
        target_id   = "...my_target_id..."
        target_type = "...my_target_type..."
      }
      repetitions = 5
    }
    medium = {
      handoff_step = {
        target_id   = "...my_target_id..."
        target_type = "...my_target_type..."
      }
      repetitions = 10
    }
  }
  repetitions   = 0
  step_strategy = "...my_step_strategy..."
  team_id       = "...my_team_id..."
}