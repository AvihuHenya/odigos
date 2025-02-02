import type { K8sActualSource } from './sources';
import type { ActualDestination } from './destinations';
import type { ActionData, ActionDataParsed } from './actions';
import type { InstrumentationRuleSpec, InstrumentationRuleSpecMapped } from './instrumentation-rules';

export type K8sActualNamespace = {
  name: string;
  selected: boolean;
  k8sActualSources?: K8sActualSource[];
};

interface PaginatedSources {
  nextPage: string;
  items: K8sActualSource[];
}

interface ComputePlatformData {
  id: string;
  name: string;
  computePlatformType: string;
  k8sActualNamespaces: K8sActualNamespace[];
  k8sActualNamespace: K8sActualNamespace;
  sources?: PaginatedSources;
  destinations: ActualDestination[];
  actions: ActionData[];
  instrumentationRules: InstrumentationRuleSpec[];
}

export type ComputePlatform = {
  computePlatform: ComputePlatformData;
};

interface ComputePlatformDataMapped extends ComputePlatformData {
  actions: ActionDataParsed[];
  instrumentationRules: InstrumentationRuleSpecMapped[];
  sources: undefined;
}

export type ComputePlatformMapped = {
  computePlatform: ComputePlatformDataMapped;
};
