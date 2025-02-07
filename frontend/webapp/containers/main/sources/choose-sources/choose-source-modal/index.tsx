import React from 'react';
import { type IAppState } from '@/store';
import { useKeyDown } from '@odigos/ui-utils';
import { useSourceCRUD, useSourceFormData } from '@/hooks';
import { ChooseSourcesBody } from '../choose-sources-body';
import { Modal, NavigationButtons } from '@odigos/ui-components';

interface Props {
  isOpen: boolean;
  onClose: () => void;
}

export const AddSourceModal: React.FC<Props> = ({ isOpen, onClose }) => {
  useKeyDown({ key: 'Enter', active: isOpen }, () => handleSubmit());

  const menuState = useSourceFormData();
  const { persistSources, loading } = useSourceCRUD({ onSuccess: onClose });

  const handleSubmit = async () => {
    const { getApiSourcesPayload, getApiFutureAppsPayload } = menuState;

    // Type of "getApiSourcesPayload()" is actually:
    // { [namespace: string]: Pick<K8sActualSource, 'name' | 'kind' | 'selected' | 'numberOfInstances'>[] };
    //
    // But we will force it as type:
    // { [namespace: string]: K8sActualSource[] };
    //
    // This forced type is to satisfy TypeScript,
    // while knowing that this doesn't break the onboarding flow in any-way...

    await persistSources(getApiSourcesPayload() as IAppState['configuredSources'], getApiFutureAppsPayload());
  };

  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      header={{ title: 'Add Source' }}
      actionComponent={
        <NavigationButtons
          buttons={[
            {
              label: 'DONE',
              variant: 'primary',
              onClick: handleSubmit,
              disabled: loading,
            },
          ]}
        />
      }
    >
      <ChooseSourcesBody componentType='FAST' isModal {...menuState} />
    </Modal>
  );
};
