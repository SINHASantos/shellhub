<template>
  <div>
    <v-list-item
      @click="showDialog = true"
      :disabled="notHasAuthorization"
      data-test="member-edit-btn"
    >
      <div class="d-flex align-center">
        <div class="mr-2">
          <v-icon> mdi-pencil </v-icon>
        </div>

        <v-list-item-title data-test="member-edit-title">
          Edit
        </v-list-item-title>
      </div>
    </v-list-item>

    <v-dialog max-width="450" v-model="showDialog">
      <v-card class="bg-v-theme-surface" data-test="member-edit-dialog">
        <v-card-title class="text-h5 pa-4 bg-primary" data-test="member-edit-dialog-title">
          Update member role
        </v-card-title>
        <v-divider />

        <v-card-text class="mt-4 mb-0 pb-1">
          <v-row align="center">
            <v-col cols="12">
              <v-select
                v-model="newRole"
                :items="items"
                label="Role"
                :error-messages="errorMessage"
                require
                data-test="role-select"
              />
            </v-col>
          </v-row>
        </v-card-text>

        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" data-test="close-btn" @click="close()">
            Close
          </v-btn>

          <v-btn
            color="primary"
            variant="text"
            data-test="edit-btn"
            @click="editMember()"
          >
            Edit
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import axios from "axios";
import { INamespaceMember } from "@/interfaces/INamespace";
import { useStore } from "@/store";
import handleError from "@/utils/handleError";
import useSnackbar from "@/helpers/snackbar";

const { member, notHasAuthorization } = defineProps({
  member: {
    type: Object as () => INamespaceMember,
    required: false,
    default: {} as INamespaceMember,
  },
  notHasAuthorization: {
    type: Boolean,
    default: false,
  },
});
const emit = defineEmits(["update"]);
const store = useStore();
const snackbar = useSnackbar();
const showDialog = ref(false);
const newRole = ref(member.role);
const errorMessage = ref("");
const items = ["administrator", "operator", "observer"];

const close = () => {
  showDialog.value = false;
};

const update = () => {
  emit("update");
  close();
};

const handleEditMemberError = (error: unknown) => {
  if (axios.isAxiosError(error)) {
    const status = error.response?.status;
    switch (status) {
      case 400:
        errorMessage.value = "The user isn't linked to the namespace.";
        break;
      case 403:
        errorMessage.value = "You don't have permission to assign a role to the user.";
        break;
      case 404:
        errorMessage.value = "The username doesn't exist.";
        break;
      default:
        handleError(error);
    }

    snackbar.showError("Failed to update user role.");
  } else {
    snackbar.showError("Failed to update user role.");
    handleError(error);
  }
};

const editMember = async () => {
  try {
    await store.dispatch("namespaces/editUser", {
      user_id: member.id,
      tenant_id: store.getters["auth/tenant"],
      role: newRole.value,
    });

    snackbar.showSuccess("Successfully updated user role.");
    update();
  } catch (error: unknown) {
    handleEditMemberError(error);
  }
};

</script>
